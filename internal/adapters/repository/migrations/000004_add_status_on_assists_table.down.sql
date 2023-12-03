-- Step 1: Remove the Default Value (if applicable)
ALTER TABLE assists
ALTER COLUMN status DROP DEFAULT;

-- Step 2: Drop the Column
ALTER TABLE assists
DROP COLUMN IF EXISTS status;